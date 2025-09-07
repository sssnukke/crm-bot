from aiogram.fsm.state import StatesGroup, State


class AddEmployeeState(StatesGroup):
    fio = State()
    phone = State()
    age = State()
    photo = State()