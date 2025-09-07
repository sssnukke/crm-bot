from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton

START_KEYBOARD = InlineKeyboardMarkup(
    inline_keyboard=[
        [InlineKeyboardButton(text='Добавить сотрудника', callback_data='add_employee')],
        [InlineKeyboardButton(text='Список сотрудников', callback_data='list_employee')],
    ]
)

def one_action_keyboard(name, callback: str) -> InlineKeyboardMarkup:
    return InlineKeyboardMarkup(
        inline_keyboard=[
            [InlineKeyboardButton(text=name, callback_data=callback)],
        ]
    )