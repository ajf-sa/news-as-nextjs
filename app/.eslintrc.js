module.exports = {
    env: {
        browser: true,
        es2021: true,
        node: true
    },
    extends: [
        "eslint:recommended",
        "plugin:@typescript-eslint/recommended",
        "plugin:react/recommended",
        'plugin:react/recommended',
        "plugin:react-hooks/recommended",
        "plugin:prettier/recommended",
    ],
    parser: '@typescript-eslint/parser',
    parserOptions: {
        ecmaFeatures: {
            jsx: true,
        },
        ecmaVersion: 12,
        sourceType: 'module',
    },
    plugins: ['react', '@typscript-eslint'],
    rules: {
        "react/react-in-jsx-scope": "off",
        "react-hooks/exhaustive-deps": "off",

        "@typescript-eslint/camelcase": "off",
        "@typescript-eslint/no-empty-function": "off",
        "@typescript-eslint/no-explicit-any": "off",
        "@typescript-eslint/ban-ts-ignore": "off",
        "@typescript-eslint/explicit-function-return-type": "off",
        "@typescript-eslint/7006": "off",
        "@typescript-eslint/explicit-module-boundary-types": "off",
    },
};