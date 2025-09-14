from aiogram import Router, F, types, Bot
from aiogram.fsm.context import FSMContext

from bot.misc import functions
from bot import keyboards
from bot.states import AddEmployeeState

callback_router = Router()


@callback_router.callback_query(F.data == 'tasks')