Smart Prompt
============

Utilitário para criar prompts bash espertos.

![Smartprompt in action](https://raw.githubusercontent.com/fholiveira/smartprompt/master/demo.gif)

##Instalar

Baixe o [binário da última versão](https://github.com/fholiveira/smartprompt/releases) e coloque-o em uma pasta que esteja no seu `$PATH`.

##Usar

Adicione a seguinte linha ao seu `~/.bashrc` para usar o prompt padrão:
```bash
export PROMPT_COMMAND='export PS1="$(smartprompt)"'´
```

Se você for usar um prompt customizado adicione:
```bash
export PROMPT_PATTERN="meu prompt customizado"
export PROMPT_COMMAND='export PS1="$(smartprompt --pattern="$PROMPT_PATTERN")"'
```

##Rodar a partir do código fonte

Este projeto depende do [git2go](https://github.com/libgit2/git2go), portanto você deve instalá-la.

Rode `go get github.com/fholiveira/smartprompt` para baixar o código e navegue para o diretório `$GOPATH/src/github.com/fholiveira/smartprompt`. Neste diretório rode `go install `para compilar e instalar os binários em `$GOPATH/bin/`. Adicione esta pasta ao seu `$PATH` ou copie os binários para uma pasta que esteja no seu `$PATH`.

##Rodar os testes
Primeiro, você deve instalar o [gorc](https://github.com/stretchr/gorc). Depois, na pasta `$GOPATH/src/github.com/fholiveira/smartprompt`, rode o comando `gorc`.

##Manual

O smartprompt, por padrão, tem o seguinte aspecto:

Este prompt default é definido pela seguinte linha:

`{GREEN:bold}{user}@{host} {BLUE:bold}{location:vimstyle} {sourcecontrol} {PURPLE:bold}{symbol} {TEXT:reset}`

Os valores envoltos em chaves são plugins. Um plugin pode aplicar uma cor ou exibir uma informação. Você pode definir um prompt diferente do padrão usando a opção `--pattern="meu prompt customizado"` ao rodar o smartprompt.

###Plugins

```
{host}            Nome do host
{user}            Nome do usuário
{dir}             Nome do diretório atual
{fqdn}            Full qualified domain name
{line:break}      Quebra de linha
{symbol}   Usa '#' quando o usuário for root e '$' para os demais usuários
{shell}           Nome do shell
{shell:version}   Versão do shell
{shell:release}   Release do shell
{location}        Caminho até o diretório atual
```
#####{symbol|*`<root>`*|*`<common user>`*}
Exibe `<root>` quando o usuário for root e `<common>` quando for um usuário comum. Se os parâmetros não forem especificados exibe "#" para root e "$" para usuário comum:

```
Usuário root:
    {symbol}        #
    {symbol|>}      >
    {symbol|>>|->}  >>

Usuário comum:
    {symbol}        $
    {symbol|>}      $
    {symbol|>>|->}  ->
```

#####{virtualenv|*`<prefix>`*|*`<sufix>`*}
Exibe o Python Virtualenv atual; os parâmetros *`<prefix>`* e *`<sufix>`* serão adicionados ao começo e fim do nome do virtualenv. Se o virtualenv for uma pasta oculta, o nome dele será mostrado sem o "." inicial.

```
{virtualenv|(|)}    (env)
{virtualenv|↦ }	    ↦ env
```

#####{location:vimstyle}
Caminho até o diretório atual usando a sintax do vim. Se diretório atual for '/mnt/pendrive/music':

```
{location:vimstyle}    /m/p/music
```

#####{time|*`<pattern>`*}
Indica a data usando o formato especificado no parâmetro *`<pattern>`*. Usando como exemplo a data '09 de Janeiro de 2014 as 21:07:02'

```
{time|dd/yy/mm}    09/01/2014
{time|d/m/y}       9/1/14
{time|h:M:s}       09:7:2
{time|hh-MM-ss}    21-07-02
```

#####{git}
Se o diretório atual for um repo git exibe as informações de acordo com o padrão **[T H S M U C]**, onde:

```
T:    status (pode ser 'merging', 'rebasing' ou vazio)
H:    head (geralmente nome da branch atual)
S:    quantidade de arquivos staged
M:    quantidade de arquivos modificados
U:    quantidade de arquivos untracked
C:    quantidade de arquivos em conflito
```
###Cores
São suportadas as cores BLACK, RED, GREEN, YELLOW, BLUE, PURPLE, CYAN, WHITE.
Elas podem ser usadas da seguinte forma:

```
{COR}               Aplica a cor ao texto
{COR:underline}     Aplica a cor ao texto e o sublinha
{COR:bold}          Aplica a cor ao texto e o coloca em negrito
{COR:background}    Aplica a cor ao background
```

Para voltar o texto para a formatação padrão use o plugin {TEXT:reset}
