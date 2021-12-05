import {Player} from '../state';
import {Action} from 'redux';

const defaultState: Player[] = [];

export default function players(
  state: Player[] = defaultState,
  _action: Action,
): Player[] {
  return state;
}
