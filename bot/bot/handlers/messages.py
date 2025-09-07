from aiogram import Router, types
from aiogram.filters import CommandStart
from aiogram.fsm.context import FSMContext

from bot import keyboards
from bot.service import users_service

messages_router = Router()


@messages_router.message(CommandStart())
async def start_handler(message: types.Message, state: FSMContext):
    await state.clear()

    user = await users_service.user_exists(message.from_user.id)
    if not user['exists']:
        await users_service.create({ "TgId": message.from_user.id })

    return message.answer('Привет, я твой CRM-bot\n'
                          'Сейчас ты можешь только создавать себе сотрудников',
                        reply_markup=keyboards.START_KEYBOARD)
