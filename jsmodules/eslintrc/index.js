// Copyright 2016 The Noms Authors. All rights reserved.
// Licensed under the Apache License, version 2.0:
// http://www.apache.org/licenses/LICENSE-2.0

module.exports = {
  'parser': 'babel-eslint',
  'rules': {
    'array-bracket-spacing': [2, 'never'],
    'arrow-body-style': [2, 'as-needed'],
    'arrow-parens': 0,
    'arrow-spacing': [2, {'before': true, 'after': true}],
    'camelcase': 2,
    'comma-dangle': [2, 'always-multiline'],
    'eol-last': 2,
    'eqeqeq': 2,
    'flow-vars/define-flow-type': 1,
    'flow-vars/use-flow-type': 1,
    'indent': [2, 2, {'SwitchCase': 1}],
    'linebreak-style': [2, 'unix'],
    'max-len': [2, 100, 8],
    'no-fallthrough': 2,
    'no-multi-spaces': 2,
    'no-new-wrappers': 2,
    'no-throw-literal': 2,
    'no-trailing-spaces': 2,
    'no-unused-vars': [2, {'argsIgnorePattern': '^_$', 'varsIgnorePattern': '^_$'}],
    'no-var': 2,
    'object-curly-spacing': [2, 'never'],
    'prefer-arrow-callback': 2,
    'prefer-const': 2,
    'quotes': [2, 'single'],
    'radix': 2,
    'semi': 2,
    'space-after-keywords': 2,
    'space-before-function-paren': 0,
    'space-in-parens': [2, 'never'],
    'space-infix-ops': 2,
  },
  'env': {
    'es6': true,
    'node': true,
    'browser': true,
  },
  'extends': 'eslint:recommended',
  'ecmaFeatures': {
    'jsx': true,
    'experimentalObjectRestSpread': true,
  },
  'plugins': [
    'flow-vars',
    'react',
  ],
};
