/* eslint-env node */
// import('@rushstack/eslint-patch/modern-module-resolution')
import js from '@eslint/js'
import vueParser from 'vue-eslint-parser'
import pluginVue from 'eslint-plugin-vue'
import pluginVueTs from '@vue/eslint-config-typescript/recommended.js'

export default [
	js.configs.recommended,
	...pluginVue.configs['vue3-recommended'],
	pluginVueTs,
	{
		files: ['src/**/*.{js,ts,vue}'],
		// root: true,
		// env: {
		// 	browser: true,
		// 	es2022: true,
		// 	node: true,
		// },
		rules: {
			'quotes': ['error', 'single'],
			'comma-dangle': ['error', 'always-multiline'],
			'semi': ['error', 'never'],

			'vue/v-on-event-hyphenation': ['warn', 'never', {'autofix': true}],
			'vue/multi-word-component-names': 'off',

			// uncategorized rules:
			'vue/component-api-style': ['error', ['script-setup']],
			'vue/component-name-in-template-casing': ['warn', 'PascalCase'],
			'vue/custom-event-name-casing': ['error', 'camelCase'],
			'vue/define-macros-order': 'error',
			'vue/match-component-file-name': ['error', {
				'extensions': ['.js', '.jsx', '.ts', '.tsx', '.vue'],
				'shouldMatchCase': true,
			}],
			'vue/no-boolean-default': ['warn', 'default-false'],
			'vue/match-component-import-name': 'error',
			'vue/prefer-separate-static-class': 'warn',

			'vue/padding-line-between-blocks': 'error',
			'vue/next-tick-style': ['error', 'promise'],
			'vue/block-lang': [
				'error',
				{'script': {'lang': 'ts'}},
			],
			'vue/no-required-prop-with-default': ['error', {'autofix': true}],
			'vue/no-duplicate-attr-inheritance': 'error',
			'vue/no-empty-component-block': 'error',
			'vue/html-indent': ['error', 'tab'],

			// vue3
			'vue/no-ref-object-destructure': 'error',
		},
		languageOptions: {
			parser: vueParser,
			ecmaVersion: 'latest',
		},
		// parserOptions: {
		// 	parser: '@typescript-eslint/parser',
		// },
		ignores: ['*.test.*'],
	}]
