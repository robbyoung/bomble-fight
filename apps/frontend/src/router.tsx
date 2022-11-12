import {BrowserRouter, Route, Link, Routes} from 'react-router-dom';
import Host from './host';
import Player from './player';

function NavMenu() {
  return (
    <div>
      <nav>
        <ul>
          <li>
            <Link to="/host">Host</Link>
          </li>
          <li>
            <Link to="/player">Player</Link>
          </li>
        </ul>
      </nav>
    </div>
  );
}

export default function Router() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<NavMenu />} />
        <Route path="/host" element={<Host />} />
        <Route path="/player" element={<Player />} />
      </Routes>
    </BrowserRouter>
  );
}
