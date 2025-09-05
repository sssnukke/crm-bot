from aiogram import Router, types
from aiogram.filters import CommandStart
from bot.service import users_service

messages_router = Router()


@messages_router.message(CommandStart())
async def start_handler(message: types.Message):

    save = await users_service.create({
        "TgId": message.from_user.id,
    })

    print(save)

    return message.answer('Привет, сохранил твой тг айди чтобы потом тебе присылать веселое что-то')
