from .callbacks import callback_router
from .messages import messages_router

tasks_router = [
    callback_router,
    messages_router,
]