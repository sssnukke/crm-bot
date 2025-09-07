from aiogram import Router, F, types, Bot
from aiogram.fsm.context import FSMContext

from bot.misc import functions
from bot import keyboards
from bot.states import AddEmployeeState

callback_router = Router()


@callback_router.callback_query(F.data == 'add_employee')
async def add_employee(callback: types.CallbackQuery, state: FSMContext):
    await state.set_state(AddEmployeeState.fio)
    await functions.delete_message(callback.bot, callback.message.chat.id, callback.message.message_id)
    message = await callback.message.answer(
        text='Напишите ФИО сотрудника\n\n'
             'Пример: Иванов Иван Иванович',
        reply_markup=keyboards.one_action_keyboard('В меню', 'menu')
    )

    await state.update_data(last_message_id=message.message_id)