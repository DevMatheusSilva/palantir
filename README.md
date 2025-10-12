# 🔮 Palantír

Palantír é uma ferramenta mágica que ajuda você a abrir seus projetos diretamente no VSCode através do terminal, inspirada no universo do Senhor dos Anéis.

## 📋 Pré-requisitos

Antes de começar, você precisará ter instalado em sua máquina:

- [Go](https://golang.org/) (versão 1.25.1 ou superior)
- [Visual Studio Code](https://code.visualstudio.com/)

## 🚀 Instalação

1. Clone o repositório:
```
bash git clone [https://github.com/DevMatheusSilva/palantir.git](https://github.com/DevMatheusSilva/palantir.git)
``` 

2. Entre no diretório do projeto:
```
bash cd palantir
``` 

3. Instale as dependências:
```
bash go mod download
``` 

4. Compile o projeto:
```
bash go build
``` 

## ⚙️ Configuração

1. Configure a variável de ambiente `PALANTIR_ROOT_FOLDER`. Esta variável deve apontar para o diretório base onde seus projetos estão localizados:

No Windows:
```
bash set PALANTIR_ROOT_FOLDER=Projects
``` 

No Linux/MacOS:
```
bash export PALANTIR_ROOT_FOLDER=Projects
``` 

> 💡 Dica: Adicione esta variável ao seu arquivo de perfil (.bashrc, .zshrc ou similar) para torná-la permanente.

## 🎯 Como Usar

O Palantír utiliza um comando simples para abrir seus projetos:
```
bash palantir open <nome_da_pasta>
``` 

Por exemplo:
```
bash palantir open meu-projeto
``` 

O Palantír irá procurar recursivamente pela pasta especificada dentro do diretório definido em `PALANTIR_ROOT_FOLDER` e abrirá o projeto no VSCode quando encontrá-lo.

## 🎭 Comportamento

- O Palantír ignora automaticamente diretórios comuns como `.git`, `node_modules` e `vendor` durante a busca
- Se o projeto não for encontrado, uma mensagem amigável será exibida
- Se o VSCode não estiver instalado, você será notificado

## 🤝 Contribuindo

Sua contribuição será muito bem-vinda! Se você tem uma ideia para melhorar o Palantír:

1. Faça um Fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/AmazingFeature`)
3. Commit suas mudanças (`git commit -m 'Add some AmazingFeature'`)
4. Push para a branch (`git push origin feature/AmazingFeature`)
5. Abra um Pull Request

## 🐛 Encontrou um bug?

Se você encontrar algum problema, por favor abra uma issue descrevendo o problema encontrado. Não se esqueça de incluir os passos para reproduzir o erro e qual comportamento você esperava.

## 📝 Licença

Este projeto está licenciado sob os termos da licença MIT.

---

*"Mesmo a menor das pessoas pode mudar o curso do futuro." - Galadriel*
