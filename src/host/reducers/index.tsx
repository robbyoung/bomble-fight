import {combineReducers} from 'redux';
import players from './players';
import combatants from './combatants';
import fight from './fight';

const rootReducer = combineReducers({fight, players, combatants});
export default rootReducer;
