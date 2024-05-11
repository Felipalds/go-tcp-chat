CREATE TRIGGER admin_also_member
    AFTER INSERT ON rooms
BEGIN
    INSERT INTO user_room(room_id, user_id, is_banned)
    VALUES(NEW.id, NEW.admin_id, false);
END;
