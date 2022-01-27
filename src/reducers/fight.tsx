import {Fight, FightStatus} from '../state';
import {Action} from 'redux';
import {ActionType} from '../actions/actionType';
import {
  progressFight,
  ProgressFightSuccessAction,
} from '../actions/progressFight';

const defaultState: Fight = {
  attackerId: '',
  defenderId: '',
  attackerDamage: 0,
  defenderDamage: 0,
  attackerHealth: 50,
  defenderHealth: 50,
  status: FightStatus.Pending,
};

export default function fight(
  state: Fight = defaultState,
  action: Action,
): Fight {
  switch (action.type) {
    case ActionType.PROGRESS_FIGHT_SUCCESS:
      return progressFight(state, action as ProgressFightSuccessAction);
    default:
      return state;
  }
}
