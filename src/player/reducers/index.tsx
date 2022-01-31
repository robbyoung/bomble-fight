import {combineReducers} from 'redux';
import player from './player';
import combatants from './combatants';
import bet from './bet';

const rootReducer = combineReducers({player, combatants, bet});
export default rootReducer;
