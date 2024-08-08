" Define syntax highlighting for account language (al)
syntax match alKey   /^[^:]\+/
syntax match alDelimiter /:/
syntax match alValue /[^:]\+$/

" Set highlight groups for the defined syntax elements
highlight def link alKey Type
highlight def link alDelimiter Statement
highlight def link alValue String

let b:current_syntax = 'foo'

" autocmd BufRead,BufNewFile *.al set filetype=al

