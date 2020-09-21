import * as React from 'react';
import * as Adapter from 'enzyme-adapter-react-16';
import * as Enzyme from 'enzyme';
import { shallow } from 'enzyme';
import { Hello } from './Hello';

Enzyme.configure({ adapter: new Adapter() });

describe('Hello.tsx', () => {
  const makeComponent = () => <Hello name="test" />;
  const wrapper = shallow(makeComponent());
  test('show Hello,', () => {
    expect(wrapper.find('div').text()).toBe('HELLO,test');
  });
});
