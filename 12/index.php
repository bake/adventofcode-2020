<?php

declare(strict_types=1);

namespace AoC\Day12;

use Generator;

class Action
{
  public function __construct(
    public string $direction,
    public int $steps,
  ) {}
}

class Point
{
  public function __construct(
    public int $x,
    public int $y,
  ) {}

  public function add(Point $p): Point
  {
    return new Point(
      $this->x + $p->x,
      $this->y + $p->y,
    );
  }

  public function multiply(int $r): Point
  {
    return new Point(
      $this->x * $r,
      $this->y * $r,
    );
  }
}

/**
 * @param handle $handle
 * @return Generator<Action> $actions
 */
function input($handle): Generator
{
  while ($line = fgets($handle)) {
    yield new Action($line[0], (int) substr($line, 1));
  }
}

/**
 * @param handle $handle
 * @return int
 */
function part1($handle): int{
  $pos = new Point(0, 0);
  $dir = new Point(1, 0);
  $input = input($handle);
  foreach ($input as $action) {
    for($a = 0; $a < $action->steps; $a += 90) {
      // I mainly just wanted an excuse to use match.
      $dir = match ($action->direction) {
        'L' => new Point($dir->y, -$dir->x),
        'R' => new Point(-$dir->y, $dir->x),
        default => $dir,
      };
    }
    $delta = match ($action->direction) {
      'N' => new Point(0, -$action->steps),
      'E' => new Point($action->steps, 0),
      'S' => new Point(0, $action->steps),
      'W' => new Point(-$action->steps, 0),
      'F' => $dir->multiply($action->steps),
      default => new Point(0,0),
    };
    $pos = $pos->add($delta);
  }
  return abs($pos->x) + abs($pos->y);
}

echo part1(STDIN);
