import React from 'react';
import jsYaml from 'js-yaml';
import TextArea from './TextArea';
import opspecDataValidator from '@opspec/sdk/lib/data/array/validator';

export default ({name, onArgChange, array}) => (
  <TextArea
    description={array.description}
    name={name}
    value={jsYaml.safeDump(array.default ? array.default : '')}
    validate={value => opspecDataValidator.validate(jsYaml.safeLoad(value), Object.assign({type: 'array'}, array.constraints))}
    onChange={value => onArgChange({array: jsYaml.safeLoad(value)})}
  />
);
