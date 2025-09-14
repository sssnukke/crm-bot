from aiogram import Router, F, types, Bot
from aiogram.fsm.context import FSMContext

from bot.misc import functions
from bot import keyboards
from bot.service import employee_service, headers

callback_router = Router()

@callback_router.callback_query(F.data == 'list_employee')
async def list_employee(callback: types.CallbackQuery, state: FSMContext):
    await state.update_data(current_page=0)
    await functions.delete_message(callback.bot, callback.message.chat.id, callback.message.message_id)


    await callback.message.answer(
        text='Сотрудники',
        reply_markup=await keyboards.list_employee.pagination_employee_keyboard(callback.from_user.id, 0)
    )


@callback_router.callback_query(F.data.startswith('list_employee-'))
async def list_employee_pagination(callback: types.CallbackQuery, state: FSMContext):
    data = await state.get_data()
    current_page = data.get('current_page', 0)

    if callback.data == 'list_employee-prev_page':
        current_page -= 1
    elif callback.data == 'list_employee-next_page':
        current_page += 1

    await state.update_data(current_page=current_page)

    try:
        await callback.message.edit_reply_markup(
            reply_markup=await keyboards.list_employee.pagination_employee_keyboard(callback.from_user.id, current_page)
        )
    except Exception:
        await functions.delete_message(callback.bot, callback.message.chat.id, callback.message.message_id)
        await callback.message.answer(
            text='Сотрудники',
            reply_markup=await keyboards.list_employee.pagination_employee_keyboard(callback.from_user.id, current_page)
        )


@callback_router.callback_query(F.data.startswith('employee_id-'))
async def employee_id(callback: types.CallbackQuery):
    await functions.delete_message(callback.bot, callback.message.chat.id, callback.message.message_id)

    employee_data = await employee_service.by_id(callback.data.split('-')[1])

    photo=types.URLInputFile(
        headers=headers,
        url=f"http://back:3000/{employee_data['photoUrl']}",
        filename=employee_data['photoUrl']
    )

    await callback.message.answer_photo(
        photo=photo,
        caption=f'ФИО: {employee_data['lastName']} {employee_data['name']} {employee_data['surName']}\n'
                f'Дата рождения: {employee_data['birthDay']}\n'
                f'Номер: {employee_data['phone']}\n',
        reply_markup=keyboards.one_action_keyboard('В меню', 'menu')
    )