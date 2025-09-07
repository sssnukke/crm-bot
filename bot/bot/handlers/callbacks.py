from aiogram import Router, F, types, Bot
from aiogram.fsm.context import FSMContext

from bot.misc import functions
from bot import keyboards

callback_router = Router()


@callback_router.callback_query(F.data == 'menu')
async def menu(callback: types.CallbackQuery, state: FSMContext):
    await state.clear()
    await functions.delete_message(callback.bot, callback.message.chat.id, callback.message.message_id)
    return callback.message.answer('Привет, я твой CRM-bot\n'
                          'Сейчас ты можешь только создавать себе сотрудников',
                          reply_markup=keyboards.START_KEYBOARD)