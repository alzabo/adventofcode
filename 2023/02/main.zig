const std = @import("std");
const info = std.log.info;
const parseInt = std.fmt.parseInt;
const expect = std.testing.expect;
const assert = std.debug.assert;

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    const in_stream = buf_reader.reader();

    var buf: [1024]u8 = undefined;
    var ct1: u32 = 0;
    var ct2: u32 = 0;

    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        const game = try parseLine(line);
        if (game.red <= 12 and game.green <= 13 and game.blue <= 14) {
            ct1 += game.id;
        }
        ct2 += game.red * game.green * game.blue;
    }

    info("part1: {d}", .{ct1});
    info("part2: {d}", .{ct2});
}

fn part1(reader: anytype) !void {
    _ = reader;
}

fn part2(reader: anytype) !void {
    var buf: [1024]u8 = undefined;
    const ct1: u32 = 0;
    _ = ct1;

    while (try reader.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        _ = line;
        //info("vals 0: {u}, 1: {u}", .{ vals[0], vals[1] });
        //info("parsing {s}...", .{vals});
        //ct1 += try parseInt(u8, &vals, 10);
    }
    //info("part2: {d}", .{ct1});
}

test "parse line" {
    const line = "Game 999: 2 green, 10 red, 3 blue; 1 blue, 5 green, 11 red; 6 red, 1 blue, 2 green; 11 red; 4 red, 1 blue, 5 green; 5 green, 3 blue";
    const v = try parseInt(u8, &[_]u8{ '0', '2' }, 10);
    _ = v;
    try parseLine(line);
}

const Game = struct {
    id: u8,
    red: u32,
    blue: u32,
    green: u32,
};

fn parseLine(line: []const u8) !Game {
    var buf = [3]u8{ 0, 0, 0 };
    var i: usize = 0;
    var val: u8 = 0;
    var game: Game = Game{ .id = 0, .red = 0, .blue = 0, .green = 0 };
    //info("buf: {s}", .{buf});
    for (line) |c| {
        switch (c) {
            '0'...'9' => {
                buf[i] = c;
                i += 1;
            },
            ':' => {
                //info("buf: {b}", .{buf});
                const end: usize = std.mem.indexOfScalar(u8, &buf, 0) orelse buf.len;
                game.id = try parseInt(u8, buf[0..end], 10);
                buf = [_]u8{ 0, 0, 0 };
                i = 0;
            },
            ' ' => {
                if (buf[0] == 0) {
                    continue;
                }
                //info("buf: {b}", .{buf});
                const end: usize = std.mem.indexOfScalar(u8, &buf, 0) orelse buf.len;
                val = try parseInt(u8, buf[0..end], 10);
                buf = [_]u8{ 0, 0, 0 };
                i = 0;
            },
            'r' => {
                if (val > game.red) {
                    game.red = val;
                }
                val = 0;
            },
            'g' => {
                if (val > game.green) {
                    game.green = val;
                }
                val = 0;
            },
            'b' => {
                if (val > game.blue) {
                    game.blue = val;
                }
                val = 0;
            },
            else => {},
        }
    }
    //info("game: {d} red: {d} green: {d} blue: {d}", .{ game.id, game.red, game.green, game.blue });
    return game;
}
