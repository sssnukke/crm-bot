from .callbacks import callback_router
from .messages import messages_router

add_employee_router = [callback_router, messages_router]