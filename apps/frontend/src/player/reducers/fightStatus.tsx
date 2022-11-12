import {FightStatus} from '../../models';
import {Action} from 'redux';
import {ActionType} from '../actions/actionType';
import {GetPlayerStateSuccessAction} from '../actions/getPlayerState';

const defaultState: FightStatus = FightStatus.Pending;

export default function fightStatus(
  state: FightStatus = defaultState,
  action: Action,
): FightStatus {
  switch (action.type) {
    case ActionType.GET_STATE_SUCCESS:
      return (action as GetPlayerStateSuccessAction).response.FightStatus;
    default:
      return state;
  }
}
