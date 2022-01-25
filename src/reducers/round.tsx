import {Round} from '../state';
import {Action} from 'redux';

const defaultState: Round = {
  combatantIds: [],
};

export default function round(
  state: Round = defaultState,
  _action: Action,
): Round {
  return state;
}
