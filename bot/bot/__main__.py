import asyncio
import logging
import sys

from aiogram import Bot

from redis.asyncio.client import Redis

from bot.misc.configuration import conf

from bot.dispatcher import get_redis_storage, get_dispatcher


async def start_bot():
    bot = Bot(conf.bot.token)
    print(conf.redis.passwd)
    print(conf.redis.host)
    storage = get_redis_storage(
        redis=Redis(
            host=conf.redis.host,
            password=conf.redis.passwd,
            port=conf.redis.port,
        )
    )
    dp = get_dispatcher(storage=storage)

    await dp.start_polling(
        bot,
        allowed_updates=dp.resolve_used_update_types(),
    )


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO, stream=sys.stdout)
    asyncio.run(start_bot())
