#users trigger start
#create
drop trigger if exists tr_users_create;
create trigger tr_users_create
    before insert
    on users
    for each row
begin
    set new.created_at = current_timestamp,
        new.active = true,
        new.updated_at = null,
        new.deleted_at = null;
end;
#update
drop trigger if exists tr_users_update;
create trigger tr_users_update
    before update
    on users
    for each row
begin
    set new.updated_at = current_timestamp;
    if old.created_at is not null then
        set new.created_at = old.created_at;
    end if;

    if new.deleted_at is not null then
        set new.active = false;
    end if;

    if old.id != new.id then
        set new.id = old.id;
    end if;
end;
#users trigger end

#companies trigger start
#create
drop trigger if exists tr_companies_create;
create trigger tr_companies_create
    before insert
    on companies
    for each row
begin

    set @inn = CAST(new.inn as CHAR(32));
    set @len = length(@inn);

    if new.inn > 0 and @len != 9 then
        signal sqlstate '45000' set message_text = 'length inn can only be 9';
    end if;

    set new.created_at = current_timestamp,
        new.active = true,
        new.updated_at = null,
        new.deleted_at = null;
end;
#update
drop trigger if exists tr_companies_update;
create trigger tr_companies_update
    before update
    on companies
    for each row
begin
    set new.updated_at = current_timestamp;

    set @inn = CAST(new.inn as CHAR(32));
    set @len = length(@inn);

    if new.inn = 0 then
        set new.inn = old.inn;
    elseif new.inn > 0 and @len != 9 then
        signal sqlstate '45000' set message_text = 'length inn can only be 9';
    end if;

    if old.created_at is not null then
        set new.created_at = old.created_at;
    end if;

    if new.deleted_at is not null then
        set new.active = false;
    end if;

    if old.id != new.id then
        set new.id = old.id;
    end if;
end;
#companies trigger end


#branches trigger start
#create
drop trigger if exists tr_branches_create;
create trigger tr_branches_create
    before insert
    on branches
    for each row
begin

    if new.until_date is null then
        set new.until_date = current_timestamp;
    end if;

    set new.created_at = current_timestamp,
        new.active = true,
        new.updated_at = null,
        new.deleted_at = null;
end;
#update
drop trigger if exists tr_branches_update;
create trigger tr_branches_update
    before update
    on branches
    for each row
begin
    set new.updated_at = current_timestamp;

    if new.until_date is null then
        set new.until_date = old.until_date;
    end if;

    if old.created_at is not null then
        set new.created_at = old.created_at;
    end if;

    if new.deleted_at is not null then
        set new.active = false;
    end if;

    if old.id != new.id then
        set new.id = old.id;
    end if;
end;
#branches trigger end
