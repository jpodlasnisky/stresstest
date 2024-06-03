## StressTest

### Parâmetros

--url ou -u (obrigatório): A URL da aplicação Web que será testada.
--requests ou -r: O número total de requisições a serem enviadas. Se esse valor não for definido, o padrão é 100.
--concurrency ou -c: Número de requisições concorrentes. Se esse valor não for definido, o padrão é 10.
--header ou -h: Cabeçalho a ser incluído na requisição. Pode ser especificado várias vezes para múltiplos cabeçalhos.

#### 1
docker build -t <nome_da_imagem> -t Dockerfile .

#### 2
docker run <nome_da_imagem> stressTest --url=http://www.pudim.com.br --requests=120 --concurrency=2
ou
docker run <nome_da_imagem> stressTest --url=http://www.pudim.com.br --requests=120 --concurrency=2 --header='API_KEY: TOKEN_1'

