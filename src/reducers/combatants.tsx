import {Combatant} from '../state';
import {Action} from 'redux';

const defaultState: Combatant[] = [];

export default function combatants(
  state: Combatant[] = defaultState,
  _action: Action,
): Combatant[] {
  return state;
}
