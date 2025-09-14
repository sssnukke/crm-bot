from .messages import messages_router
from .callbacks import callback_router
from .add_employee import add_employee_router
from .list_employee import list_employee_router
from .tasks import tasks_router

routers = (
    messages_router,
    callback_router,
    *add_employee_router,
    *list_employee_router,
    *tasks_router,
)
