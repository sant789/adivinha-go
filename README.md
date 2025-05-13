# Jogo de Adivinhação 🎲

<br>

## 🇧🇷 Versão em Português (pt-BR)

<div align="center">
  
![Go Version](https://img.shields.io/badge/Go-1.23.3-00ADD8?style=for-the-badge&logo=go)
![Bubble Tea](https://img.shields.io/badge/Bubble_Tea-TUI-7D56F4?style=for-the-badge)

**Um jogo de adivinhação no terminal com uma interface elegante e interativa!**
</div>

### 📝 Sobre o Projeto
Este projeto foi desenvolvido como parte da minha jornada de aprendizado em Go (Golang). É um jogo simples de adivinhação onde você tenta adivinhar um número entre 0 e 100 dentro de 10 tentativas.

> **Nota:** Este código foi construído durante um período de estudos e pesquisa sobre Golang, explorando conceitos de desenvolvimento de aplicações TUI (Terminal User Interface) e a utilização de bibliotecas modernas do ecossistema Go.

### ✨ Funcionalidades
- Interface de linha de comando (TUI) interativa e colorida
- Dicas visuais baseadas na proximidade ao número secreto ("quente" ou "frio")
- Barra de progresso para acompanhar tentativas restantes
- Histórico visual dos palpites anteriores
- Feedback em tempo real sobre cada tentativa

### 🛠️ Tecnologias Utilizadas
- [Go 1.23.3](https://go.dev/) - Linguagem de programação
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Framework para criação de TUIs
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Biblioteca para estilização de texto
- [Bubbles](https://github.com/charmbracelet/bubbles) - Componentes reutilizáveis para Bubble Tea

### 🚀 Como Executar

#### Pré-requisitos
- Go 1.23.3 ou superior instalado

#### Instalação
```bash
# Clone o repositório
git clone https://github.com/Paulo-Borszcz/go-1---Guesing-Game.git
cd adivinha-go

# Baixe as dependências
go mod tidy

# Execute o jogo
go run main.go
```

### 🎯 Como Jogar
1. O jogo escolherá um número aleatório entre 0 e 100
2. Digite seu palpite e pressione Enter
3. O jogo informará se seu palpite está acima ou abaixo do número secreto
4. Você terá no máximo 10 tentativas para adivinhar o número
5. Dicas "quente" ou "frio" indicarão quão perto você está

### 📚 O Que Aprendi
- Criação de aplicações TUI (Terminal User Interface) com Go
- Uso do padrão Model-View-Update (MVU) do Bubble Tea
- Manipulação de entrada de usuário e eventos de teclado
- Estilização de texto no terminal com Lip Gloss
- Gerenciamento de estados em aplicações Go
- Uso de estruturas de dados e funções em Go
- Implementação de lógica de jogo simples

### 🔍 O Que Poderia Ser Melhorado
Como este é um projeto de estudo, há várias melhorias que poderiam ser implementadas:
- Adicionar níveis de dificuldade
- Implementar um sistema de pontuação
- Adicionar sons ou animações mais elaboradas
- Salvar estatísticas de jogos anteriores
- Implementar testes automatizados

---
# Guessing Game 🎲

## 🇺🇸 English Version (en-US)

<div align="center">
  
![Go Version](https://img.shields.io/badge/Go-1.23.3-00ADD8?style=for-the-badge&logo=go)
![Bubble Tea](https://img.shields.io/badge/Bubble_Tea-TUI-7D56F4?style=for-the-badge)

**A guessing game in the terminal with an elegant and interactive interface!**
</div>

### 📝 About the Project
This project was developed as part of my learning journey in Go (Golang). It's a simple guessing game where you try to guess a random number between 0 and 100 within 10 attempts.

> **Note:** This code was built during a period of study and research on Golang, exploring concepts of TUI (Terminal User Interface) application development and the use of modern libraries from the Go ecosystem.

### ✨ Features
- Interactive and colorful command-line interface (TUI)
- Visual hints based on proximity to the secret number ("hot" or "cold")
- Progress bar to track remaining attempts
- Visual history of previous guesses
- Real-time feedback on each attempt

### 🛠️ Technologies Used
- [Go 1.23.3](https://go.dev/) - Programming language
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Framework for creating TUIs
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Library for text styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - Reusable components for Bubble Tea

### 🚀 How to Run

#### Prerequisites
- Go 1.23.3 or higher installed

#### Installation
```bash
# Clone the repository
git clone https://github.com/Paulo-Borszcz/go-1---Guesing-Game.git
cd guessing-go

# Download dependencies
go mod tidy

# Run the game
go run main.go
```

### 🎯 How to Play
1. The game will choose a random number between 0 and 100
2. Type your guess and press Enter
3. The game will tell you if your guess is above or below the secret number
4. You will have a maximum of 10 attempts to guess the number
5. "Hot" or "cold" hints will indicate how close you are

### 📚 What I Learned
- Creating TUI (Terminal User Interface) applications with Go
- Using the Model-View-Update (MVU) pattern of Bubble Tea
- Handling user input and keyboard events
- Styling text in the terminal with Lip Gloss
- Managing state in Go applications
- Using data structures and functions in Go
- Implementing simple game logic

### 🔍 What Could Be Improved
As this is a learning project, there are several improvements that could be implemented:
- Add difficulty levels
- Implement a scoring system
- Add sounds or more elaborate animations
- Save statistics from previous games
- Implement automated tests
