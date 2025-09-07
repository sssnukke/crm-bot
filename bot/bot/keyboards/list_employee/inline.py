import math

from aiogram.types import InlineKeyboardMarkup, InlineKeyboardButton

from bot.service import users_service


async def pagination_employee_keyboard(tg_id: int, current_page) -> InlineKeyboardMarkup:
    employees = await users_service.user_by_tg_id(tg_id)

    total_pages = math.ceil(len(employees['employees']) / 6)

    start_index = current_page * 6
    end_index = start_index + 6
    end_index = min(end_index, len(employees['employees']))

    page_items = employees['employees'][start_index:end_index]

    buttons = []
    for employee in page_items:
        buttons.append([InlineKeyboardButton(text=f'{employee["lastName"]} {employee["name"]}', callback_data=f'employee_id-{employee["ID"]}')])

    prev_callback_data = f'list_employee-prev_page' if current_page > 0 else '#'
    next_callback_data = f'list_employee-next_page' if current_page < total_pages - 1 else '#'

    if len(employees['employees']) > 5:
        navigation_buttons = [
            InlineKeyboardButton(text='⬅️', callback_data=prev_callback_data),
            InlineKeyboardButton(text=f'{current_page + 1}/{total_pages}', callback_data='#'),
            InlineKeyboardButton(text='➡️', callback_data=next_callback_data)
        ]

        buttons.append(navigation_buttons)

    buttons.append([InlineKeyboardButton(text='Вернуться в меню', callback_data='menu')])

    return InlineKeyboardMarkup(inline_keyboard=buttons)