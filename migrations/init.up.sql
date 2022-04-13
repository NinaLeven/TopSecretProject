create table "project"
(
    uid             varchar(36)  not null primary key,
    name            varchar(512) not null,
    owner_id        varchar(512) not null,
    participant_ids varchar(512)[] not null,
    progress        smallint     not null,
    state           varchar(100) not null,
    created_at      timestamp    not null,
    updated_at      timestamp    not null
);

-- indices are not strictly needed; i decided to show my ability to work with them

create index "project_state_index" on "project" using btree("state");