# 先决条件

## From Github Releases

[just](https://github.com/casey/just/releases/tag/1.25.2)

## Linux

### Alpine linux

```bash
apk add just
```

### Arch linux

```bash
pacman -S just
```

### Debian and Ubuntu 衍生品

```bash
git clone https://mpr.makedeb.org/just
cd just
makedeb -si
```
## Windows

### Winget

```powershell
winget install --id Casey.Just --exact
```

### Scoop

```powershell
scoop install just
```

### Chocolatey

```powershell
choco install just
```