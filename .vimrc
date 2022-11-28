"vim-plug start
call plug#begin('~/.config/nvim/plugged')
Plug 'preservim/nerdtree'
Plug 'ryanoasis/vim-devicons'
Plug 'tomasiser/vim-code-dark'
Plug 'vim-airline/vim-airline'
Plug 'vim-airline/vim-airline-themes'
Plug 'fatih/vim-go'
call plug#end()
"vim-plug end
"vim-go doc popup
let g:go_doc_popup_window = 1
let g:go_doc_keywordprg_enable = 0
" Setting
set number
set title
set showmatch
set ruler
" 구문강조
if has("syntax")
	syntax on
endif
"색 설정
set t_Co=256
set t_ut=
colorscheme codedark
"들여쓰기 설정
set autoindent
set smartindent
set tabstop=4
set shiftwidth=4
set softtabstop=4
set smarttab
set expandtab
"붙여 넣기 설정
set paste
set mouse-=a
"한글 입력 설정
set encoding=utf-8
set termencoding=utf-8
"상태바 표시
set laststatus=2
set statusline=\ %<%l:%v\ [%P]%=%a\ %h%m%r\ %F\
"Ctrl+ n 에 NERDTree 토글
map <c-n> :NERDTreeToggle<CR>
"vim 수행 시 파일 미지정 또는 디렉터리 지정 시 NERDTree 수행되도록 설정
autocmd StdinReadPre * let s:std_in=1
autocmd VimEnter * if argc() == 0 && !exists("s:std_in") | NERDTree | endif
autocmd VimEnter * if argc() == 1 && isdirectory(argv()[0]) && !exists("s:std_in") | exe 'NERDTree' argv()[0] | wincmd p | ene | exe 'cd '.argv()[0] | endif
" 마지막으로 수정된 곳에 커서를 위치함
au BufReadPost *
\ if line("'\"") > 0 && line("'\"") <= line("$") |
\ exe "norm g`\"" |
\ endif
" 스크롤바 너비
set sw=1
