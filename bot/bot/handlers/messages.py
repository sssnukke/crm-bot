from aiogram import Router, types
from aiogram.filters import CommandStart

messages_router = Router()


@messages_router.message(CommandStart())
async def start_handler(message: types.Message):
    return message.answer('Hi')
