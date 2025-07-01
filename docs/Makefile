.PHONY: default clean hugo hugo-build

default: hugo

clean:
	rm -rf public/

hugo-build: clean hugo-themes
	hugo --enableGitInfo --source .

hugo:
	# hugo server --disableFastRender --enableGitInfo --watch --source .
	hugo server -D -E -F
	# hugo server -D

hugo-themes:
	[ -d themes ] && rm -rf themes
	[ -d themes ] || mkdir -pv themes
	git clone --depth=1 https://github.com/alex-shpak/hugo-book themes/hugo-book
	rm -rf themes/hugo-book/.git

hugo-themes-save:
	rm -rf themes
	mkdir themes
	# git clone --depth=1 https://github.com/thingsym/hugo-theme-techdoc.git themes/hugo-theme-techdoc
	# rm -rf themes/hugo-theme-techdoc/.git
	git clone --depth=1 https://github.com/alex-shpak/hugo-book themes/hugo-book
	rm -rf themes/hugo-book/.git
	# git clone --depth=1 https://github.com/docura/docura themes/docura
	# rm -rf themes/docura/.git
	# git clone --depth=1 https://github.com/bowman2001/perplex themes/hugo-perplex
	# rm -rf themes/hugo-perplex/.git
	# git clone --depth=1 https://github.com/adityatelange/hugo-PaperMod themes/hugo-PaperMod
	# rm -rf themes/hugo-PaperMod/.git
