export enum ActionType {
  GET_PLAYERS_REQUEST = 'Get Players Request',
  GET_PLAYERS_SUCCESS = 'Get Players Success',
  GET_PLAYERS_FAILURE = 'Get Players Failure',

  START_PLAYERS_POLL = 'Start Polling Players',
  STOP_PLAYERS_POLL = 'Stop Polling Players',

  GET_COMBATANTS_REQUEST = 'Get Combatants Request',
  GET_COMBATANTS_SUCCESS = 'Get Combatants Success',
  GET_COMBATANTS_FAILURE = 'Get Combatants Failure',

  PROGRESS_FIGHT_REQUEST = 'Progress Fight Request',
  PROGRESS_FIGHT_SUCCESS = 'Progress Fight Success',
  PROGRESS_FIGHT_FAILURE = 'Progress Fight Failure',

  RESET_FIGHT_REQUEST = 'Reset Fight Request',
  RESET_FIGHT_SUCCESS = 'Reset Fight Success',
  RESET_FIGHT_FAILURE = 'Reset Fight Failure',
}
