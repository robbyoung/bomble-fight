import {combineReducers} from 'redux';
import player from './player';
import combatants from './combatants';
import bet from './bet';
import fightStatus from './fightStatus';

const rootReducer = combineReducers({player, combatants, bet, fightStatus});
export default rootReducer;
