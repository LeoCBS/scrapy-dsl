# scrapy-dsl (aka sdsl)

Reading the book [Writing an interpreter in Go](https://interpreterbook.com/) i
decided to build one very simple DSL to study and have fun :)

## Proposal

DSL to simplify write spiders. Propose this domain specific language is
abstract python and scrapy knowledgement and provide one language more simple to write spiders.

## DSL syntax proposal
    
1)     
 
    get 'http://www.telelistas.net' extract phone on xpath('/div[@id=telInfo]');


2) 
    let firtBody = get('http://www.telelistas.net')
    let results = parser(firstBody){
        return append({"fone": xpath('/div[@id=telInfo]')})
    }

After running this command scrapy-dsl will performe one request to source and
extract field fone where xpath match.

Output Example:

    {"fone": "12345-1245"}, 
    {"fone": "12345-0000"}, 
    {"fone": "48000-1245"}, 
    {"fone": "47000-1245"}

## How install

## How run unit tests

    make check

## REPL

Like other languages SDSL also have one REPL (Read, Eval, Print, Loop). This
mode also can called to "console" mode or "interactive" mode.

To interact with SDSL run this command:

    go run main.go

