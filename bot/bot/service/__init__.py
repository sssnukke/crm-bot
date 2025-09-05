from ..misc.configuration import conf

url = 'http://back:3000'

headers = {
    "Authorization": f"Bearer {conf.bot.secret_token}"
}