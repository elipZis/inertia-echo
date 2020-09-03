const mix = require('laravel-mix');
const cssImport = require('postcss-import');
const cssNesting = require('postcss-nesting');
const path = require('path');
const purgecss = require('@fullhuman/postcss-purgecss');
const tailwindcss = require('tailwindcss');

require('laravel-mix-svelte');

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel applications. By default, we are compiling the CSS
 | file for the application as well as bundling up all the JS files.
 |
 */

mix.js('resources/js/app.js', 'js')
    .postCss('resources/css/app.css', 'css/app.css')
    .svelte({
        dev: !mix.inProduction()
    })
    .options({
        postCss: [
            cssImport(),
            cssNesting(),
            tailwindcss('tailwind.config.js'),
            ...(mix.inProduction()
                ? [
                    purgecss({
                        content: [
                            './resources/views/**/*.html',
                            './resources/js/**/*.svelte'
                        ],
                        defaultExtractor: content =>
                            content.match(/[\w-/:.]+(?<!:)/g) || [],
                        whitelistPatternsChildren: [/nprogress/]
                    })
                ]
                : [])
        ]
    })
    .webpackConfig({
        output: {chunkFilename: 'js/[name].js?id=[chunkhash]'},
        resolve: {
            alias: {
                '@': path.resolve('resources/js')
            }
        }
    })
    .setPublicPath('./public')
    .version()
    .sourceMaps();
