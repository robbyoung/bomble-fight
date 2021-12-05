import {combineReducers} from 'redux';
import game from './game';
import players from './players';
import combatants from './combatants';
import round from './round';

const rootReducer = combineReducers({game, round, players, combatants});
export default rootReducer;
