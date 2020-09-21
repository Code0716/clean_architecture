import * as React from 'react';
import { useEffect } from 'react';
import { AutoSizer, Column, Table } from 'react-virtualized';
import { imageListActions } from '../../actions/imageListActions';
import { Portal } from '../../components/Portal';
import { FormContainer } from '../../components/FormContainer';
import './style.scss';

export const ListImages = () => {
  const {
    //state
    images,
    //action
    forwordToDetaile,
    deletePreupload,
  } = imageListActions();

  return (
    <AutoSizer>
      {({ width }) => (
        <React.Fragment>
          <Table
            data={images}
            height={800}
            width={width}
            headerHeight={35}
            rowHeight={35}
            rowGetter={({ index }) => images[index]}
            rowCount={images.length}
            rowClassName="virtualized_row"
            onRowClick={({ rowData }) => forwordToDetaile(rowData.info.ID)}
          >
            <Column
              width={100}
              label="ID"
              dataKey=""
              className="colum_row"
              cellDataGetter={({ rowData }) => rowData.info.ID}
            />
            <Column
              flexGrow={1}
              width={200}
              label="Images"
              dataKey=""
              className="colum_row"
              cellDataGetter={({ rowData }) => rowData.info.Title}
            />
            <Column
              flexGrow={1}
              width={200}
              label="Path"
              dataKey=""
              className="colum_row"
              cellDataGetter={({ rowData }) => rowData.info.Path}
            />
          </Table>
          {images.length !== 0 && (
            <Portal domId={'upload'} className="">
              <div className="image_commit_box">
                <div>Upload</div>
                <button
                  className="action_button"
                  onClick={() => deletePreupload('upload')}
                >
                  <span>Delete all images</span>
                </button>
              </div>
            </Portal>
          )}
        </React.Fragment>
      )}
    </AutoSizer>
  );
};
