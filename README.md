# 🔮 Palantír - The Seeing Stone

![Palantír - The Seeing Stone](https://raw.githubusercontent.com/DevMatheusSilva/palantir/main/.github/assets/palantir.jpg)

*"A great seeing-stone to aid in your quest through the realms of code."*

Palantír is a mystical tool forged in the depths of Golang, empowering you to swiftly open your projects in VSCode through the terminal, much like the ancient seeing-stones of Middle-earth guided the wise.

## � Prerequisites

Before embarking on your journey, ensure you have these artifacts in your possession:

- 🧙‍♂️ [Go](https://golang.org/) (version 1.25.1 or higher) - The language of power
- ⚔️ [Visual Studio Code](https://code.visualstudio.com/) - Your trusty weapon in the battle against bugs

## 🌟 Installation

### The Path of the Wise

1. First, summon the repository from the depths of GitHub:
```bash
git clone https://github.com/DevMatheusSilva/palantir.git
```

2. Enter the sacred grounds:
```bash
cd palantir
```

3. Gather the ancient dependencies:
```bash
go mod download
```

4. Install the seeing-stone in your realm (this will place the executable in your Go workspace):
```bash
go install
```

5. Create a symbolic link to wield its power from anywhere (requires sudo):
```bash
sudo ln -s $GOPATH/bin/palantir /usr/local/bin/palantir
```

## 🗝️ Configuration

Just as the Palantíri needed proper alignment to function, you must configure the sacred path:

1. Set the `PALANTIR_ROOT_FOLDER` environment variable to point to your projects' realm:

For the Kingdoms of Men (Windows):
```bash
set PALANTIR_ROOT_FOLDER=Projects
```

For the Realms of Elves and Dwarves (Linux/MacOS):
```bash
export PALANTIR_ROOT_FOLDER=Projects
```

> 🧙‍♂️ **Gandalf's Wisdom**: Add this variable to your profile scroll (.bashrc, .zshrc, or similar) to make it permanent.

## ⚔️ How to Wield the Seeing Stone

The Palantír responds to a simple incantation:
```bash
palantir open <project-name>
```

For instance, to peer into the realm of your project:
```bash
palantir open my-project
```

Like the ancient seeing-stones, Palantír will search recursively through the realms defined in `PALANTIR_ROOT_FOLDER` and open your chosen project in VSCode when found.

## 🎭 The Magic Within

- Much like Gandalf avoiding the paths of the enemy, Palantír automatically ignores common directories like `.git`, `node_modules`, and `vendor`
- If your project remains hidden, a friendly message shall guide you
- Should VSCode be missing from your realm, you shall be notified

## 🤝 Join the Fellowship

![The Fellowship](https://raw.githubusercontent.com/DevMatheusSilva/palantir/main/.github/assets/fellowship.jpg)

Your aid in improving the Palantír shall be most welcome! To join our fellowship:

1. Fork the sacred repository
2. Create a branch for your quest (`git checkout -b feature/AmazingFeature`)
3. Commit your enchantments (`git commit -m 'Add some AmazingFeature'`)
4. Push to the realm (`git push origin feature/AmazingFeature`)
5. Open a Pull Request to share your wisdom

## 🐛 Encountered an Orc (Bug)?

Should you encounter any dark forces (bugs) in your journey, please open an issue describing your tale. Remember to include the steps to summon the error and what outcome you expected, so our fellowship may aid you in vanquishing it.

## � License

This artifact is licensed under the terms of the MIT scroll.

---

![Middle-earth Map](https://raw.githubusercontent.com/DevMatheusSilva/palantir/main/.github/assets/map.jpg)

*"Even the smallest person can change the course of the future." - Galadriel*
