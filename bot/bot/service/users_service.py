import aiohttp

from . import url, headers

async def create(data: dict) -> dict | None:
    async with aiohttp.ClientSession(
            headers=headers
    ) as session:
        try:
            return await (await session.post(
                f'{url}/users',
                json=data
            )).json()
        except aiohttp.client_exceptions.ContentTypeError:
            return None
