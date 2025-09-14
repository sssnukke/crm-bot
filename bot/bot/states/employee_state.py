from aiogram.fsm.state import StatesGroup, State


class AddEmployeeState(StatesGroup):
    fio = State()
    phone = State()
    birthday = State()
    photo = State()