
# Como executar:
  - 1) Criar uma instância do mysql no docker: docker run -p 3306:3306 --name golang_sql -e MYSQL_ROOT_PASSWORD=root -d mysql:8.0.29 ou rodar na maquina na porta 3306
  - 2) git clone https://github.com/marcoscoutinhodev/golang_api.git
  - 3) cd golang_api
  - 4) go run ./main.go (Instância MYSQL precisa esta rodando na porta 3306 quando startar o server para criar o DB e a Tabela de Task)

# Rotas:
  - Adicionar nova task: /POST http://127.0.0.1:4000/task/add | Enviar JSON informando o 'title: string' e 'description: string'  (O campo done é automaticamente salvo como false)
  - Buscar todas as tasks: /GET http://127.0.0.1:4000/task/get
  - Buscar task por id: /GET http://127.0.0.1:4000/task/get-by-id?id=<TASK_ID> | Informar o ID na query string
  - Atualizar task por id: /POST http://127.0.0.1:4000/task/update?id=<TASK_ID> | Informar o ID na query string e Enviar JSON informando PELO MENOS 1 campo para atualizar (title: string, description: string, done: boolean)
  - Deletar task por id: /POST http://127.0.0.1:4000/task/delete?id=<TASK_ID> | Informar o ID na query string
  - Alternar o campo 'done' de uma task por id: /POST http://127.0.0.1:4000/task/toggle-done-field?id=<TASK_ID> | Informar o ID na query string
  
# Video:
  - https://www.youtube.com/watch?v=TNHz5B4kTWQ