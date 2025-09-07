import re
import base64

from aiogram import Router, types
from aiogram.fsm.context import FSMContext

from bot import keyboards
from bot.misc import functions
from bot.states import AddEmployeeState
from bot.service import employee_service

messages_router = Router()


@messages_router.message(AddEmployeeState.fio)
async def add_fio_employee(message: types.Message, state: FSMContext):
    data = await state.get_data()

    await functions.delete_message(message.bot, message.chat.id, data['last_message_id'])
    await functions.delete_message(message.bot, message.chat.id, message.message_id)

    fio_pattern = re.compile(r'^[А-ЯЁ][а-яё]+\s[А-ЯЁ][а-яё]+\s[А-ЯЁ][а-яё]+$')

    if message.text and fio_pattern.match(message.text):
        await state.set_state(AddEmployeeState.age)
        await state.update_data(fio=message.text)
        message = await message.answer(
            text='Напишите возраст сотрудника',
            reply_markup=keyboards.one_action_keyboard('Вернуться в меню','menu')
        )
    else:
        message = await message.answer(
            text='Ошибка попробуйте еще раз,\n\n'
                 'Пример: Иванов Иван Иванович',
            reply_markup=keyboards.one_action_keyboard('Вернуться в меню', 'menu')
        )

    await state.update_data(last_message_id=message.message_id)


@messages_router.message(AddEmployeeState.age)
async def add_age_employee(message: types.Message, state: FSMContext):
    data = await state.get_data()

    await functions.delete_message(message.bot, message.chat.id, data['last_message_id'])
    await functions.delete_message(message.bot, message.chat.id, message.message_id)

    if message.text and message.text.strip().isdigit():
        await state.set_state(AddEmployeeState.phone)
        await state.update_data(age=message.text)
        message = await message.answer(
            text='Напишите номер телефона сотрудника',
            reply_markup=keyboards.one_action_keyboard('Вернуться в меню', 'menu')
        )
    else:
        message = await message.answer(
            text='Ошибка попробуйте еще раз,\n\n'
                 'Пример: 24',
            reply_markup=keyboards.one_action_keyboard('Вернуться в меню', 'menu')
        )

    await state.update_data(last_message_id=message.message_id)


@messages_router.message(AddEmployeeState.phone)
async def add_phone_employee(message: types.Message, state: FSMContext):
    data = await state.get_data()

    await functions.delete_message(message.bot, message.chat.id, data['last_message_id'])
    await functions.delete_message(message.bot, message.chat.id, message.message_id)

    phone_pattern = re.compile(r'^\+7\d{10}$')

    if message.text and phone_pattern.match(message.text):
        await state.set_state(AddEmployeeState.photo)
        await state.update_data(phone=message.text)
        message = await message.answer(
            text='Отправьте фотографию сотрудника',
            reply_markup=keyboards.one_action_keyboard('Вернуться в меню', 'menu')
        )
    else:
        message = await message.answer(
            text='Ошибка попробуйте еще раз,\n\n'
                 'Пример: +79144200499',
            reply_markup=keyboards.one_action_keyboard('Вернуться в меню', 'menu')
        )

    await state.update_data(last_message_id=message.message_id)


@messages_router.message(AddEmployeeState.photo)
async def add_photo_employee(message: types.Message, state: FSMContext):
    data = await state.get_data()

    await functions.delete_message(message.bot, message.chat.id, data['last_message_id'])
    await functions.delete_message(message.bot, message.chat.id, message.message_id)

    if message.photo:
        photo = message.photo[-1]
        photo_file = await message.bot.download(photo)
        photo_bytes = photo_file.read()
        photo_base64 = base64.b64encode(photo_bytes).decode('utf-8')

        photo_data_url = f"data:image/jpeg;base64,{photo_base64}"

        name = data['fio'].split()[1]
        lastName = data['fio'].split()[0]
        surName = data['fio'].split()[2]

        await employee_service.create({
            'userId': message.from_user.id,
            'employee': {
                'name': name,
                'lastName': lastName,
                'surName': surName,
                'age': int(data['age']),
                'phone': data['phone'],
                'photo': photo_data_url
            }
        })

        await message.answer(
            text='Сотрудник добавлен в базу данных',
            reply_markup=keyboards.one_action_keyboard('Вернуться в меню', 'menu')
        )

    else:
        message = await message.answer(
            text='Ошибка попробуйте еще раз',
            reply_markup=keyboards.one_action_keyboard('Вернуться в меню', 'menu')
        )

        await state.update_data(last_message_id=message.message_id)

