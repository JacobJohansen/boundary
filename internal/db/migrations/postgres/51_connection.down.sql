begin;

  drop table session_connection_state;
  drop table session_connection_state_enm;
  drop table session_connection;
  drop table session_connection_closed_reason_enm;
  drop function insert_session_connection_state;

commit;
