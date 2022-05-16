import {Fight, FightAction, FightStatus} from '../../models';
import {Action} from 'redux';
import {ActionType} from '../actions/actionType';
import {
  progressFight,
  ProgressFightSuccessAction,
} from '../actions/progressFight';

const defaultState: Fight = {
  left: {
    id: '',
    health: 50,
    loss: 0,
    action: FightAction.Nothing,
  },
  right: {
    id: '',
    health: 50,
    loss: 0,
    action: FightAction.Nothing,
  },
  status: FightStatus.Pending,
};

export default function fight(
  state: Fight = defaultState,
  action: Action,
): Fight {
  switch (action.type) {
    case ActionType.PROGRESS_FIGHT_SUCCESS:
      return progressFight(state, action as ProgressFightSuccessAction);
    case ActionType.RESET_FIGHT_SUCCESS:
      return defaultState;
    default:
      return state;
  }
}
