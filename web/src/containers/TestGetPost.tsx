import * as React from 'react';
import { useState } from 'react';
import { AutoSizer, Column, Table } from 'react-virtualized';
import { Select } from '../components/Select';
import { FormContainer } from '../components/FormContainer';
import { testGetPostActions } from '../actions/testGetPostActions';
import { shotMessageAction } from '../actions/shotMessageActions';
type tableStyle = {
  width: number;
  height: number;
  index: number;
};
const TestGetPost = () => {
  const [getInput, setGetInput] = useState('/city');
  const [postInput, setPostInput] = useState('');
  const [textArea, setTextArea] = useState('');

  const { getData, postData, get, post, migrate } = testGetPostActions();
  const { showMessage } = shotMessageAction();

  const selctData: { value: string }[] = [
    { value: '/country' },
    { value: '/country/japan' },
    { value: '/city' },
    { value: '/city/tokyo' },
    { value: '/city/kyoto' },
  ];

  const disabled: boolean = true;

  return (
    <FormContainer>
      <div className="test_center mb-30">
        <div>非同期通信の確認画面</div>
        {postData}
        <h1>inputサンプル</h1>
        <label>GET</label>
        <Select
          value={getInput}
          option={selctData}
          onChange={(e: React.ChangeEvent<HTMLSelectElement>) =>
            setGetInput(e.target.value)
          }
        />
        <button className="action_button" onClick={() => get(getInput)}>
          <span>GET button</span>
        </button>
        {/*
        <label>POST(TODO)</label>
        <input
          value={postInput}
          type="text"
          onChange={() => {}}
          disabled={disabled}
        />
        <textarea
          className="test_textarea"
          disabled={disabled}
          value={textArea}
          onChange={e => setTextArea(e.target.value)}
        />
        <button
          className="action_button"
          onClick={() => post(postInput, textArea)}
          disabled={disabled}
        >
          <span>POST button</span>
        </button>*/}
      </div>
      <div className="test_center mb-30">
        <label>Notification</label>
        <div className="d-flex mb-15">
          <button
            className="action_button mr-5"
            onClick={() => showMessage('test')}
          >
            test
          </button>
          <button className="action_button" onClick={() => showMessage('rest')}>
            rest
          </button>
        </div>
      </div>
      <div className="test_center mb-30">
        <label>migration</label>
        <button className="action_button" onClick={migrate}>
          migrate
        </button>
      </div>
      {/*  <AutoSizer>
        {({ width, height }: tableStyle) => (
          <Table
            data={getData}
            height={height}
            width={width}
            rowHeight={35}
            rowGetter={({ index }: tableStyle) => getData[index]}
            rowCount={getData.length}
            rowClassName="d-flex"
          >
            <Column width={200} flexGrow={1} label="name" dataKey="name" />
            <Column width={200} label="code" dataKey="code" />
          </Table>
        )}
        </AutoSizer>*/}
    </FormContainer>
  );
};

export default TestGetPost;
