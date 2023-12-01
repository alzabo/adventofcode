const std = @import("std");
const info = std.log.info;
const parseInt = std.fmt.parseInt;

pub fn main() !void {
    var file = try std.fs.cwd().openFile("input.txt", .{});
    defer file.close();

    var buf_reader = std.io.bufferedReader(file.reader());
    var in_stream = buf_reader.reader();
    var buf: [1024]u8 = undefined;

    var ct1: u32 = 0;

    while (try in_stream.readUntilDelimiterOrEof(&buf, '\n')) |line| {
        var vals: [2]u8 = [_]u8{ 0, 0 };
        for (line) |c| {
            switch (c) {
                '0'...'9' => {
                    if (vals[0] == 0) {
                        vals[0] = c;
                    }
                    vals[1] = c;
                },
                else => {},
            }
        }
        //info("parsing {s}...", .{vals});
        ct1 += try parseInt(u8, &vals, 10);
    }
    info("part1: {d}", .{ct1});
}
