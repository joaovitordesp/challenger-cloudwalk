[Explicação arquitetural]

- cmd/parser/main.go: Esse arquivo será responsável por inicializar o parser e passar o caminho do arquivo de log para a função de parsing.
- pkg/logparser/parser.go: Contém a lógica principal para ler o arquivo de log e agrupar os dados dos jogos.
- pkg/logparser/models.go: Definirá as estruturas de dados que representarão as informações dos jogos, jogadores e kills.
- tests/test.go: Testes unitários para garantir que o parser funcione corretamente.
- logs/quake_log.txt: O arquivo de log do Quake 3 Arena que será analisado.

- Como rodar o programa
  - na raiz do projeto, rode o comando
    - go run cmd/parser/main.go
