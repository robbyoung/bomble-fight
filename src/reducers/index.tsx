import {combineReducers} from 'redux';
import game from './game';
import players from './players';
import combatants from './combatants';
import fight from './fight';
import bets from './bets';

const rootReducer = combineReducers({game, fight, players, combatants, bets});
export default rootReducer;
