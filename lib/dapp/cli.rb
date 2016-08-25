module Dapp
  # CLI
  class CLI
    extend Helper::Cli
    include Mixlib::CLI
    include Helper::Trivia

    SUBCOMMANDS = %w(build push spush list run stages cleanup bp).freeze

    banner <<BANNER.freeze
Usage: dapp [options] sub-command [sub-command options]

Available subcommands: (for details, dapp SUB-COMMAND --help)

dapp build [options] [APPS PATTERN ...]
dapp bp [options] [APPS PATTERN ...] REPO
dapp push [options] [APP PATTERN] REPO
dapp spush [options] [APPS PATTERN ...] REPO
dapp list [options] [APPS PATTERN ...]
dapp run [options] [APP PATTERN] [DOCKER ARGS]
dapp cleanup [options] [APPS PATTERN ...]
dapp stages

Options:
BANNER

    option :version,
           long: '--version',
           description: 'Show version',
           on: :tail,
           boolean: true,
           proc: proc { puts "dapp: #{Dapp::VERSION}" },
           exit: 0

    option :help,
           short: '-h',
           long: '--help',
           description: 'Show this message',
           on: :tail,
           boolean: true,
           show_options: true,
           exit: 0

    def initialize(*args)
      super(*args)

      opt_parser.program_name = 'dapp'
      opt_parser.version = Dapp::VERSION
    end

    def run(argv = ARGV)
      argv, subcommand, subcommand_argv = self.class.parse_subcommand(self, argv)
      self.class.parse_options(self, argv)
      self.class.run_subcommand self, subcommand, subcommand_argv
    end
  end
end
