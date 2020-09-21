// action type
export const ActionTypes = {
  LOADING: 'LOADING',
  UNLOADING: 'UNLOADING',
} as const;

// action creator
/**
 * 処理中
 */
export function loading() {
  return {
    type: ActionTypes.LOADING,
  };
}

/**
 * 処理後
 */
export function unloading() {
  return {
    type: ActionTypes.UNLOADING,
  };
}
