import {combineReducers} from 'redux';
import game from './game';
import players from './players';
import combatants from './combatants';
import round from './round';
import bets from './bets';

const rootReducer = combineReducers({game, round, players, combatants, bets});
export default rootReducer;
