export enum ActionType {
  ADD_PLAYER_REQUEST = 'Add Player Request',
  ADD_PLAYER_SUCCESS = 'Add Player Success',
  ADD_PLAYER_FAILURE = 'Add Player Failure',

  GET_STATE_SUCCESS = 'Get Player State Success',
  GET_STATE_FAILURE = 'Get Player State Failure',

  START_STATE_POLL = 'Start Polling Player State',
  STOP_STATE_POLL = 'Stop Polling Player State',

  GET_COMBATANTS_REQUEST = 'Get Combatants Request',
  GET_COMBATANTS_SUCCESS = 'Get Combatants Success',
  GET_COMBATANTS_FAILURE = 'Get Combatants Failure',

  ADD_BET_REQUEST = 'Add Bet Request',
  ADD_BET_SUCCESS = 'Add Bet Success',
  ADD_BET_FAILURE = 'Add Bet Failure',

  PROGRESS_FIGHT_REQUEST = 'Progress Fight Request',
  PROGRESS_FIGHT_SUCCESS = 'Progress Fight Success',
  PROGRESS_FIGHT_FAILURE = 'Progress Fight Failure',
}
